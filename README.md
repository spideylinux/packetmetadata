## packetmetadata

This is a thin golang SDK and CLI for interacting with Packet's gRPC metadata service. 


### library usage

```golang
import "github.com/packethost/packetmetadata/packetmetadata"

res, err := packetmetadata.Get()


iterator, err := packetmetadata.Watch()
for {
  next, err := iterator.Next()
  fmt.Println(string(next.Patch), string(next.JSON))
}
```


### cli usage
```
docker run quay.io/packet/packetmetadata watch
docker run quay.io/packet/packetmetadata get
```
