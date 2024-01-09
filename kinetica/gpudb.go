package kinetica

import (
	"bytes"
	"context"
	"crypto/tls"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"github.com/golang/snappy"
	"github.com/hamba/avro/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"

	"github.com/ztrue/tracerr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var (
	logger               *zap.Logger
	maxSize              = 10
	maxBackups           = 5
	maxAge               = 10
	defaultLogConfigFile = "./kinetica/log_config.yaml"
)

type Kinetica struct {
	url            string
	options        *KineticaOptions
	client         *resty.Client
	tracer         trace.Tracer
	kineticaLogger *zap.Logger
}

type KineticaOptions struct {
	Username           string
	Password           string
	UseSnappy          bool
	ByPassSslCertCheck bool
	TraceHTTP          bool
	Timeout            string
	LogConfigFile      string
}

var (
	//go:embed avsc/*
	schemaFS embed.FS
)

func New(ctx context.Context, url string) *Kinetica {
	return NewWithOptions(ctx, url, &KineticaOptions{LogConfigFile: defaultLogConfigFile})
}

func NewWithOptions(ctx context.Context, url string, options *KineticaOptions) *Kinetica {
	var (
		// childCtx context.Context
		childSpan trace.Span
	)

	client := resty.New()
	tracer := otel.GetTracerProvider().Tracer("kinetica-golang-api")
	client.DisableWarn = true

	if options.ByPassSslCertCheck {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	_, childSpan = tracer.Start(ctx, "NewWithOptions()")
	defer childSpan.End()

	if options.Timeout != "" {
		duration, err := time.ParseDuration(options.Timeout)
		if err != nil {
			fmt.Println("Error parsing timout, no timeout set.", err)
		} else {
			client.SetTimeout(duration)
		}
	}

	kinetica := &Kinetica{url: url, options: options, client: client, tracer: tracer}
	kinetica.kineticaLogger = kinetica.createLogger(options)

	return kinetica
}

func parseNumber(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err == nil {
		return v
	}
	return fallback
}

// createLogger - creates a Logger using user defined configuration from a file.
// If the file name is not found or there is some error it will create a default
// logger using the default logger config file named 'config_log_zap.yaml'.
//
//	@receiver kinetica
//	@param options
//	@return *zap.Logger
func (kinetica *Kinetica) createLogger(options *KineticaOptions) *zap.Logger {

	if logger != nil {
		return logger
	}

	if options.LogConfigFile == "" {
		return kinetica.createDefaultLogger()
	}

	var logConfig zap.Config

	yamlFile, _ := os.ReadFile(options.LogConfigFile)
	if err := yaml.Unmarshal(yamlFile, &logConfig); err != nil {
		fmt.Println("Error in loading log config file : ", options.LogConfigFile)
		fmt.Println(err)
		fmt.Println("Creating default logger ...")
		logger = kinetica.createDefaultLogger()
		return logger
	}
	fmt.Println("Log Config from YAML file : ", logConfig)

	var (
		stdout      zapcore.WriteSyncer
		file        zapcore.WriteSyncer
		logFilePath *url.URL
	)

	for _, path := range logConfig.OutputPaths {

		if path == "stdout" || path == "stderror" {
			stdout = zapcore.AddSync(os.Stdout)
			fmt.Println("Created stdout syncer ...")

		} else if strings.HasPrefix(path, "lumberjack://") {
			var err error
			logFilePath, err = url.Parse(path)
			fmt.Println("LogFilePath : ", logFilePath)

			if err == nil {
				filename := strings.TrimLeft(logFilePath.Path, "/")
				fmt.Println("LogFileName : ", filename)

				if filename != "" {
					q := logFilePath.Query()
					l := &lumberjack.Logger{
						Filename:   filename,
						MaxSize:    parseNumber(q.Get("maxSize"), maxSize),
						MaxAge:     parseNumber(q.Get("maxAge"), maxAge),
						MaxBackups: parseNumber(q.Get("maxBackups"), maxBackups),
						LocalTime:  false,
						Compress:   false,
					}

					file = zapcore.AddSync(l)
					fmt.Println("Created file syncer ...")
				}
			}
		} else {
			// Unknown output format
			fmt.Println("Invalid output path specified in config ...")
		}
	}

	if stdout == nil && file == nil {
		fmt.Println("Both stdout and file not available; creating default logger ...")
		return kinetica.createDefaultLogger()
	}

	level := zap.NewAtomicLevelAt(logConfig.Level.Level())

	consoleEncoder := zapcore.NewConsoleEncoder(logConfig.EncoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(logConfig.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)

	logger = zap.New(core)
	return logger
}

// createDefaultLogger //
//
//	@receiver kinetica
//	@param options
//	@return *zap.Logger
func (kinetica *Kinetica) createDefaultLogger() *zap.Logger {
	if logger != nil {
		return logger
	}

	var logConfig zap.Config

	yamlFile, rerr := os.ReadFile(defaultLogConfigFile)
	if rerr != nil {
		fmt.Println(rerr.Error())
		return nil
	}

	if err := yaml.Unmarshal(yamlFile, &logConfig); err != nil {
		fmt.Println("Error in loading default log config file : ", defaultLogConfigFile)
		fmt.Println(err)
		fmt.Println("No log configured ...")
		return nil
	}
	fmt.Println("Default Log Config from YAML file : ", logConfig)

	var (
		stdout      zapcore.WriteSyncer
		file        zapcore.WriteSyncer
		logFilePath *url.URL
	)

	for _, path := range logConfig.OutputPaths {

		if path == "stdout" || path == "stderror" {
			stdout = zapcore.AddSync(os.Stdout)
			fmt.Println("Created stdout syncer ...")

		} else if strings.HasPrefix(path, "lumberjack://") {
			var err error
			logFilePath, err = url.Parse(path)
			fmt.Println("LogFilePath : ", logFilePath)

			if err == nil {
				filename := strings.TrimLeft(logFilePath.Path, "/")
				fmt.Println("LogFileName : ", filename)

				if filename != "" {
					q := logFilePath.Query()
					l := &lumberjack.Logger{
						Filename:   filename,
						MaxSize:    parseNumber(q.Get("maxSize"), maxSize),
						MaxAge:     parseNumber(q.Get("maxAge"), maxAge),
						MaxBackups: parseNumber(q.Get("maxBackups"), maxBackups),
						LocalTime:  false,
						Compress:   false,
					}

					file = zapcore.AddSync(l)
					fmt.Println("Created file syncer ...")
				}
			}
		} else {
			// Unknown output format
			fmt.Println("Invalid output path specified in config ...")
		}
	}

	level := zap.NewAtomicLevelAt(logConfig.Level.Level())

	consoleEncoder := zapcore.NewConsoleEncoder(logConfig.EncoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(logConfig.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)

	kinetica.kineticaLogger = zap.New(core)
	return kinetica.kineticaLogger
}

func parseSchema(asset string) *avro.Schema {
	var (
		err          error
		schemaByte   []byte
		schemaString string
		schema       avro.Schema
	)

	schemaByte, err = schemaFS.ReadFile(asset)
	if err != nil {
		panic(err)
	}

	schemaString = strings.ReplaceAll(string(schemaByte), "\n", "")
	schema, err = avro.Parse(schemaString)
	if err != nil {
		panic(err)
	}

	return &schema
}

// GetOptions
//
//	@receiver kinetica
//	@return *GpudbOptions
func (kinetica *Kinetica) GetOptions() *KineticaOptions {
	return kinetica.options
}

func (kinetica *Kinetica) submitRawRequest(
	ctx context.Context,
	restUrl string, requestSchema *avro.Schema, responseSchema *avro.Schema,
	requestStruct interface{}, responseStruct interface{}) error {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.submitRawRequest()")
	defer childSpan.End()

	requestBody, err := avro.Marshal(*requestSchema, requestStruct)
	if err != nil {
		err = tracerr.Wrap(err)
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	httpResponse, err := kinetica.buildHTTPRequest(childCtx, &requestBody).Post(kinetica.url + restUrl)
	if err != nil {
		err = tracerr.Wrap(err)
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	childSpan.AddEvent("Response Info:", trace.WithAttributes(
		attribute.Int("Status Code: ", httpResponse.StatusCode()),
		attribute.String("Proto: ", httpResponse.Proto()),
		attribute.String("Time: ", httpResponse.Time().String()),
		attribute.String("Received At: ", httpResponse.ReceivedAt().String()),
	))

	ti := httpResponse.Request.TraceInfo()
	childSpan.AddEvent("Request Trace Info:", trace.WithAttributes(
		attribute.String("DNSLookup: ", ti.DNSLookup.String()),
		attribute.String("ConnTime: ", ti.ConnTime.String()),
		attribute.String("TCPConnTime: ", ti.TCPConnTime.String()),
		attribute.String("TLSHandshake: ", ti.TLSHandshake.String()),
		attribute.String("ServerTime: ", ti.ServerTime.String()),
		attribute.String("ResponseTime: ", ti.ResponseTime.String()),
		attribute.String("TotalTime: ", ti.TotalTime.String()),
		attribute.Bool("IsConnReused: ", ti.IsConnReused),
		attribute.Bool("IsConnWasIdle: ", ti.IsConnWasIdle),
		attribute.String("ConnIdleTime: ", ti.ConnIdleTime.String()),
		attribute.Int("RequestAttempt: ", ti.RequestAttempt),
		// attribute.String("RemoteAddr: ", ti.RemoteAddr.String()),
	))

	body := httpResponse.Body()
	bodyStr := string(body[:])
	fmt.Println(bodyStr)
	reader := avro.NewReader(bytes.NewBuffer(body), len(body))

	status := reader.ReadString()
	if reader.Error != nil {
		err = errors.New(reader.Error.Error())
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	message := reader.ReadString()
	if reader.Error != nil {
		err = errors.New(reader.Error.Error())
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	childSpan.AddEvent("Response Message:", trace.WithAttributes(
		attribute.String("Message: ", message),
	))

	if status == "ERROR" {
		err = errors.New(message)
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	reader.SkipString()
	reader.SkipInt()
	reader.ReadVal(*responseSchema, &responseStruct)
	if reader.Error != nil {
		err = errors.New(reader.Error.Error())
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, err.Error())

		return err
	}

	childSpan.AddEvent("Response Struct:", trace.WithAttributes(
		attribute.String("Spew: ", spew.Sdump(responseStruct)),
	))

	return nil
}

func (kinetica *Kinetica) buildHTTPRequest(ctx context.Context, requestBody *[]byte) *resty.Request {

	var (
	// childCtx context.Context
	// childSpan trace.Span
	)

	request := kinetica.client.R().
		SetBasicAuth(kinetica.options.Username, kinetica.options.Password)

	if kinetica.options.UseSnappy {
		snappyRequestBody := snappy.Encode(nil, *requestBody)

		request = request.SetHeader("Content-type", "application/x-snappy").SetBody(snappyRequestBody)
	} else {
		request = request.SetHeader("Content-type", "application/octet-stream").SetBody(*requestBody)
	}

	if kinetica.options.TraceHTTP {
		request = request.EnableTrace()
	}

	return request
}
