# Documentation: STI

Currently, the storage interface supports the following drivers:

- Minio     (S3 object storage)
- VFS       (Virtual Filesystem)
- Mockup    (Internal testing store)

### Driver: minio

Toggle: `-driver minio`

Environment:
- S3HOST    (Address of the S3 API)
- S3USER    (User/login name)
- S3SECRET  (Password/Secret)
- S3BUCKET  (Default target bucket)

### Driver: vfs

Toggle: `-driver vfs`

Environment:
- VFSDIR    (Default directory)

### Driver: mockup

Toggle: `-driver mockup`
