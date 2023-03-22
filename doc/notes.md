# EXM Notes

This document contains notes for engineers and developers.
EXM listens to every address on port 1800.

### Environment variables

- S3HOST    (Address of the S3 API)
- S3USER    (User/login name)
- S3SECRET  (Password/Secret)
- S3BUCKET  (Default target bucket)

- VFSDIR    (Default directory)

### Exit codes

0. Exit without errors.
1. Exit with errors (should not happen in production).

### FProbes

In EXM, every function has a `fprobe` constant.
This constant is used by the `con.Probeln()` function, and gives insights into the program flow.

The fprobe format should be: `packageName.functionName`
