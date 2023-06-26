Usage:

```bash
./class-version-tracker.exe track -g org.apache.dubbo -a dubbo -c org.apache.dubbo.rpc.protocol.DelegateExporterMap -o output
./class-version-tracker.exe track -g org.apache.dubbo -a dubbo -c org.apache.dubbo.rpc.protocol.AbstractProtocol -o output
./class-version-tracker.exe track -g org.apache.dubbo -a dubbo -c org.apache.dubbo.rpc.protocol.dubbo.DubboProtocol -o output
./class-version-tracker.exe track -g org.apache.dubbo -a dubbo -c org.apache.dubbo.common.extension.ExtensionLoader -o output

./class-version-tracker.exe track -g com.alibaba -a dubbo -c com.alibaba.dubbo.rpc.protocol.AbstractProtocol -o output
./class-version-tracker.exe track -g com.alibaba -a dubbo -c com.alibaba.dubbo.common.extension.ExtensionLoader -o output
./class-version-tracker.exe track -g com.alibaba -a dubbo -c com.alibaba.dubbo.rpc.protocol.dubbo.DubboProtocol -o output
```

# TODO

- 支持指定是使用本地mvn配置的镜像仓库还是直接连中央仓库，因为某些镜像仓库有点坑，会有那种版本空洞，冷不丁给你来一下 

