##### zocli exit codes

| Return code | Description                                                                                                                |
| ----------- | -------------------------------------------------------------------------------------------------------------------------- |
| 0           | The command execution succeeded.                                                                                           |
| 1           | The command execution failed with a completion code that signals an error.                                                 |
| 2           | The CLI was able to send a request message. However, it did not get a response message within the specified time interval. |
| 3           | The CLI was not able to send a request message.                                                                            |
| 4           | The initialization of the CLI failed.                                                                                      |
| 5           | Error related to external cli package like, Package: survey, cli/browser                                                   |
| 6           | Error related to parsing flags supplied on the cli                                                                         |
| 7           | Error in connection to server                                                                                              |
| 8           | Server returned error                                                                                                      |
| 9           | Response parsing failure                                                                                                   |
| 10          | File System Related errors                                                                                                 |
