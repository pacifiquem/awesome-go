## Log Package for Go

The log package provides a simple logging facility for Go programs. It allows you to log messages at different levels of severity, such as DEBUG, INFO, WARNING, ERROR, and FATAL. The messages are then sent to a configured output, such as the standard error stream or a file.

### Creating a Logger

To create a logger, you can use the `log.New()` function:

```go
logger := log.New(os.Stderr, "", log.LstdFlags)
```

This will create a new logger that logs to the standard error stream and uses the default flags. The flags specify the format of the log messages.

### Logging Messages

To log a message, you can use one of the following methods:

```go
logger.Debug("Starting the program")
logger.Info("Processing data")
logger.Warn("A potential error occurred")
logger.Error("An error occurred")
logger.Fatal("The program is terminating")
```

* `Debug(format string, a ...interface{})`: Logs a DEBUG-level message.
* `Info(format string, a ...interface{})`: Logs an INFO-level message.
* `Warn(format string, a ...interface{})`: Logs a WARNING-level message.
* `Error(format string, a ...interface{})`: Logs an ERROR-level message.
* `Fatal(format string, a ...interface{})`: Logs a FATAL-level message and then exits the program.

The format string is a template for the log message, and the `a ...interface{}` parameter is a list of values to be inserted into the template.

### Customizing the Logger

You can customize the logger by setting the following flags:

```go
logger.Flags() |= log.Ldate | log.Ltime | log.Lmicroseconds
logger.SetOutput(os.Stdout)
```

* `log.Ldate`: Logs the date and time of the message.
* `log.Ltime`: Logs the time of the message.
* `log.Lmicroseconds`: Logs the time of the message with microsecond precision.
* `log.Llongfile`: Logs the full filename of the source code line where the log message was created.
* `log.Lshortfile`: Logs the base name and line number of the source code line where the log message was created.
* `log.LUTC`: Logs the time in UTC instead of the local time zone.

You can also set the output of the logger to a different destination, such as a file, using the `SetOutput()` method.

### Examples


```go
logger := log.New(os.Stdout, "my-prefix", log.LstdFlags|log.Lshortfile)
logger.Debug("Starting the program")
logger.Info("Processing data")
logger.Warn("A potential error occurred")
logger.Error("An error occurred")
logger.Fatal("The program is terminating")
```

Here are some examples of how to use the log package:

### More Information

For more information about the log package, please refer to the documentation: https://pkg.go.dev/log
