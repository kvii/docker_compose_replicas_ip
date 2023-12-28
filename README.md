# IP 问题复现

1. 执行 `make image` 构建镜像。
2. 执行 `docker compose up -d`，启动工程。
3. 执行 `curl http://localhost`，观察响应的 name 和 hostname。
4. 编辑 compose.yaml，将 app1 和 app2 的 tag 改成 0.0.2，replicas 改成 5。
5. 执行 `docker compose up -d` 更新工程。
6. 再次执行 `curl http://localhost`，响应的 name 属性变为 app2。说明请求被 nginx 转发到了错误的容器上。

因为是偶发现象，所以响应的 name 可能不会变。可以再次将 app1 和 app2 的 tag 改成 0.0.3，replicas 改为 15 后再尝试一次。