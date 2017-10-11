#### 相关介绍

1. 该日志模块是基于`go.uber.org/zap`和`gopkg.in/natefinch/lumberjack.v2`这两个库
2. zap用于打印json格式的日志，lumberjack用于切割日志。
3. 需要在app.conf文件中添加两个字段:`serverName`和`logFilePath`
4. 在其他地方使用时，直接调用vlogs.VesyncLog.xxx()