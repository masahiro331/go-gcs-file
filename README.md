

## Quick Start

```go

const (
    BucketName = "trivy-vm-images"
    ObjectName = "trivy-snapshot-image.vmdk"
)

func main() {
    ctx := context.Background()
    const blockSize  = 4096
    f, err := NewFile(ctx, BucketName, ObjectName)
    if err != nil {
        log.Fatal(err)
    }
    
    buf := make([]byte, blockSize)
    n, err := f.ReadAt(buf, 0)
    if err != nil {
        log.Fatal(err)
    }
    if n != blockSize {
        log.Fatal(xerrors.New("read bytes size error"))
    }
    
    fmt.Println(hex.Dump(buf))
}
```