/*
Package pocketlog exposes an API to log your work.

First, instantiate a logger with pocketlog.New, passing it a threshold log level.
Messages of lesser criticality will not be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on three levels of criticality:
  - Debug: used to log messages for debugging code during development.
  - Info: used to log general information about the program's execution.
  - Error: used to log errors that occur during execution.
*/
package pocketlog
