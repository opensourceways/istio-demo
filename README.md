# istio-demo

## 1.这个代码仓是做什么的

是一个demo服务，基于istio服务网格，解决微服务互相调用可观测性的问题，展示微服务调用拓扑图

参考 文档：https://istio.io/latest/zh/docs/tasks/observability/kiali/

## 2.使用方法

### 1.准备一个k8s集群

我用的是mac，可以用docker-desktop在本地搭建一个k8s集群

在 macOS 上搭建 Kubernetes (K8s) 集群，你可以选择多种方法。最常见的方法是使用 **Docker Desktop** 来本地运行 Kubernetes，它提供了简单的集成和易于管理的界面。以下是如何使用 Docker Desktop 来在 macOS 上搭建和配置 Kubernetes 的详细步骤。

#### 步骤 1：安装 Docker Desktop

1. **下载 Docker Desktop for Mac**

   前往 Docker 官方网站 下载适用于 macOS 的 Docker Desktop 安装包。

2. **安装 Docker Desktop**

   下载完成后，打开 `.dmg` 文件并按照安装向导进行安装。安装完成后，启动 Docker Desktop 应用。

3. **启动 Docker Desktop**

   启动 Docker Desktop 后，它会在后台运行，并且在 Mac 的菜单栏中显示 Docker 图标。

   如果是首次启动，Docker Desktop 可能需要一些时间来进行初始化。

#### 步骤 2：启用 Kubernetes

1. **启用 Kubernetes**

   打开 Docker Desktop 应用，点击菜单栏上的 Docker 图标，选择 **Preferences**（偏好设置）。

   在 **Preferences** 窗口中，选择 **Kubernetes** 标签页。

   然后勾选 **Enable Kubernetes** 选项。启用 Kubernetes 后，Docker 会自动在后台配置 Kubernetes，并创建一个单节点的 Kubernetes 集群。

2. **配置 Kubernetes 设置**

   - 在同一窗口中，你可以选择启用或禁用一些选项，比如是否启用 `kubectl` 自动安装，是否使用虚拟机等。
   - **Memory**: 为了获得更好的性能，建议分配至少 4GB 内存。

3. **应用配置**

   配置完成后，点击 **Apply & Restart** 按钮。Docker Desktop 将会重新启动并开始启动 Kubernetes 集群，这可能需要几分钟的时间。

4. **验证 Kubernetes 是否启动**

   Kubernetes 启动后，你可以在 Docker Desktop 的 **Kubernetes** 标签页中查看集群的状态。

#### 步骤 3：安装 kubectl

1. **安装 kubectl**

   `kubectl` 是与 Kubernetes 集群进行交互的命令行工具。如果你已经安装了 Docker Desktop，那么 `kubectl` 已经被集成到 Docker Desktop 中，并且会自动配置好 Kubernetes 的 `kubeconfig` 文件。

   如果你还没有安装 `kubectl`，你可以使用 Homebrew 来安装它：

   ```
   bash
   
   
   复制代码
   brew install kubectl
   ```

2. **验证 kubectl 是否配置正确**

   在终端中执行以下命令来验证 `kubectl` 是否可以正确连接到 Kubernetes 集群：

   ```
   bash
   
   
   复制代码
   kubectl version --client
   ```

   这应该输出 `kubectl` 客户端的版本信息。

   然后，执行以下命令检查 Kubernetes 集群的状态：

   ```
   bash
   
   
   复制代码
   kubectl cluster-info
   ```

   如果集群已正确启动，你应该看到类似以下的输出：

   ```
   bash
   
   
   复制代码
   Kubernetes master is running at https://localhost:6443
   KubeDNS is running at https://localhost:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
   ```

#### 步骤 4：测试 Kubernetes 集群

1. **查看集群节点**

   使用以下命令查看当前 Kubernetes 集群的节点信息：

   ```
   bash
   
   
   复制代码
   kubectl get nodes
   ```

   输出应该显示集群中只有一个节点（因为你使用的是 Docker Desktop 本地单节点集群）：

   ```
   复制代码
   NAME             STATUS   ROLES    AGE   VERSION
   docker-desktop   Ready    master   10m   v1.21.0
   ```

2. **部署一个简单的 Pod**

   为了验证集群是否正常工作，你可以部署一个简单的 Pod（比如一个 Nginx 服务）：

   ```
   bash
   
   
   复制代码
   kubectl run nginx --image=nginx --restart=Never
   ```

3. **查看 Pod 状态**

   使用以下命令检查 Pod 是否已经成功运行：

   ```
   bash
   
   
   复制代码
   kubectl get pods
   ```

   你应该会看到一个名为 `nginx` 的 Pod 正在运行。

4. **访问应用**

   你可以通过 `kubectl port-forward` 命令将服务端口映射到本地端口：

   ```
   bash
   
   
   复制代码
   kubectl port-forward pod/nginx 8080:80
   ```

   然后，你可以在浏览器中访问 `http://localhost:8080` 来查看 Nginx 页面。

#### 步骤 5：管理和监控 Kubernetes 集群

- **Docker Desktop Dashboard**
  Docker Desktop 提供了一个可视化的 Kubernetes Dashboard，你可以在应用程序中查看 Pods、Deployments、Services 等 Kubernetes 资源。

- **查看集群资源**
  你可以使用 `kubectl` 命令行工具来查看和管理集群中的资源。例如：

  - 查看 Pod：

    ```
    bash
    
    
    复制代码
    kubectl get pods
    ```

  - 查看 Service：

    ```
    bash
    
    
    复制代码
    kubectl get svc
    ```

- **删除 Kubernetes 集群**
  如果你需要删除 Docker Desktop 上的 Kubernetes 集群，可以通过 Docker Desktop 的 **Preferences** 中的 **Kubernetes** 标签页，点击 **Disable Kubernetes** 来禁用并删除集群。

#### 总结

使用 **Docker Desktop** 是在 macOS 上搭建 Kubernetes 集群最简单和直接的方法。它提供了与 Docker 的无缝集成，允许你轻松管理本地的 Kubernetes 集群。通过这些步骤，你可以快速设置一个本地的单节点 Kubernetes 集群并开始进行开发和测试。

### 2.部署demo服务到k8s

#### go

以go-istio-client为例

1. 打包docker镜像

   ```
    docker build -t go-istio-client:<tag> .
   ```

2. 修改 [go-istio-client-deployment.yaml](/Users/shishupei/go/istio-demo/go-istio-client/go-istio-client-deployment.yaml) ，image的tag 20250109v3，与镜像tag保持一致

   ```
   spec:
     containers:
       - name: go-istio-client
         image: go-istio-client:20250109v3
         ports:
           - containerPort: 8081
   ```

  3.部署服务到k8s

```
kubectl apply -f java-istio-server-deployment.yaml
```

  4.验证，

```
kubectl get svc
```

返回正在运行的service

```
NAME                TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)           AGE
go-istio-client     LoadBalancer   10.111.230.8     localhost     30002:31325/TCP   2d21h
```

#### java

1.打jar包

```
mvn clean package
```

2.和go的步骤一致

### 安装istio

#### 步骤 1：下载 Istio 二进制文件

1. **访问 Istio 官方发布页面：**

   打开浏览器，访问 Istio 的 GitHub 发布页面： https://github.com/istio/istio/releases

2. **下载适合 macOS 的最新版本：**

   在页面中，找到最新的稳定版本并下载适用于 macOS 的二进制文件（`.tar.gz` 文件）。例如，选择一个像 `istio-1.16.1-osx.tar.gz` 这样的文件。

3. **使用 curl 下载二进制文件（也可以直接下载）：**

   你可以使用以下命令来下载最新版本的 Istio（确保在终端中执行）：

   ```
   bash
   
   
   复制代码
   curl -L https://github.com/istio/istio/releases/download/1.16.1/istio-1.16.1-osx.tar.gz -o istio-1.16.1-osx.tar.gz
   ```

   请注意替换 URL 为最新版本的下载链接。

#### 步骤 2：解压文件

1. **解压下载的文件：**

   下载完成后，使用以下命令解压 `.tar.gz` 文件：

   ```
   bash
   
   
   复制代码
   tar -xvzf istio-1.16.1-osx.tar.gz
   ```

   这会在当前目录下创建一个名为 `istio-1.16.1` 的文件夹（其中 `1.16.1` 为实际的版本号）。

#### 步骤 3：将 `istioctl` 添加到 PATH

1. **将 Istio 二进制文件添加到 PATH：**

   进入解压后的目录并将 `istioctl` 命令添加到你的环境变量中。可以通过以下命令来完成：

   ```
   bash
   
   
   复制代码
   cd istio-1.16.1
   sudo mv bin/istioctl /usr/local/bin/
   ```

   这会将 `istioctl` 命令移动到 `/usr/local/bin/`，该目录通常已包含在 PATH 中。

2. **验证 `istioctl` 是否安装成功：**

   执行以下命令来验证是否能够正确使用 `istioctl`：

   ```
   bash
   
   
   复制代码
   istioctl version
   ```

   如果安装成功，你应该看到类似以下的输出：

   ```
   yaml
   
   
   复制代码
   istioctl version: 1.16.1
   ```

#### 步骤 4：安装 Istio 到 Kubernetes 集群

现在你已经安装了 Istio CLI，可以使用 `istioctl` 来安装 Istio 控制平面和相关组件到 Kubernetes 集群中。

1. **使用 `istioctl` 安装 Istio：**

   通过以下命令将 Istio 安装到 Kubernetes 集群中：

   ```
   bash
   
   
   复制代码
   istioctl install --set profile=default
   ```

   这将使用默认的配置文件安装 Istio。你可以根据需要选择不同的配置文件（如 `minimal`, `demo`, `empty` 等）。

2. **验证 Istio 控制面是否成功安装：**

   安装完成后，运行以下命令检查 Istio 的组件（如 `istiod` 和 `istio-ingressgateway`）是否成功启动：

   ```
   bash
   
   
   复制代码
   kubectl get pods -n istio-system
   ```

   如果安装成功，你应该看到类似以下的输出：

   ```
   sql
   
   
   复制代码
   NAME                                       READY   STATUS    RESTARTS   AGE
   istio-ingressgateway-5d4f56ff7d-abcde      1/1     Running   0          5m
   istiod-7dbd5dddc9-abcdef                  1/1     Running   0          5m
   ```

#### 步骤 5：启用自动注入 Sidecar

为了让 Istio 自动注入 Sidecar，你需要为相关的命名空间添加标签。例如，如果你的应用在 `default` 命名空间下运行，可以运行以下命令：

```
bash


复制代码
kubectl label namespace default istio-injection=enabled
```

#### 步骤6:安装kiali

Kiali 是一个开源的 Istio 服务网格监控和可视化工具，它为 Istio 提供了图形化的界面，帮助你更好地理解和管理 Istio 服务网格的流量、健康状态和配置。

以下是在 Kubernetes 中安装 Kiali 的步骤：

Kiali 可以通过 Istio 提供的安装脚本来安装。你可以使用 `istioctl` 来为 Istio 安装 Kiali。

1. **安装 Kiali 使用 Istioctl：**

   Istio 提供了 Kiali 的安装配置文件，你可以通过 `istioctl` 安装 Kiali。

   执行以下命令安装 Kiali：

   ```
   bash
   
   
   复制代码
   istioctl install --set profile=demo
   ```

   这里使用了 `demo` 配置文件，这会自动安装 Istio 控制平面和 Kiali。`demo` 配置文件是一个启用了一些 Istio 插件和组件（包括 Kiali、Prometheus、Grafana 等）的配置。

2. **手动安装 Kiali（可选）：**

   如果你不想使用 `istioctl` 安装，或者需要更详细的控制，下面是使用 `kubectl` 手动安装 Kiali 的方法。

   创建一个 Kiali 安装配置文件：

   ```
   bash
   
   
   复制代码
   kubectl apply -f https://raw.githubusercontent.com/kiali/kiali/master/deploy/kubernetes/istio/kiali.yaml
   ```

   这将部署 Kiali 的必要资源，包括 Deployment、Service 和相关的 RBAC 配置。

3. **验证 Kiali 安装：**

   安装完成后，运行以下命令检查 Kiali 的 Pod 是否已经启动：

   ```
   bash
   
   
   复制代码
   kubectl get pods -n istio-system
   ```

   你应该能看到一个类似下面的 Pod：

   ```
   sql
   
   
   复制代码
   NAME                                 READY   STATUS    RESTARTS   AGE
   kiali-xyz123456-abcde               1/1     Running   0          2m
   ```

 4,**访问 Kiali UI**

Kiali 提供了一个 Web UI 用于查看和管理 Istio 服务网格的状态。你可以通过以下步骤访问它：

1. **获取 Istio Ingress Gateway 的外部 IP 地址：**

   如果你使用的是 Kubernetes 集群，Kiali 通常会通过 Istio 的 Ingress Gateway 暴露。首先，你需要获取 Ingress Gateway 的外部 IP 地址。

   运行以下命令来查看 Istio 的 Ingress Gateway 服务：

   ```
   bash
   
   
   复制代码
   kubectl get svc istio-ingressgateway -n istio-system
   ```

   输出会包含一个 `EXTERNAL-IP` 字段，类似于：

   ```
   scss
   
   
   复制代码
   NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                                    AGE
   istio-ingressgateway    LoadBalancer   10.110.240.152   <external-ip>    15021:30222/TCP,80:31380/TCP,443:31390/TCP   5m
   ```

   记下 `EXTERNAL-IP`。

2. **通过端口转发访问 Kiali UI：**

   如果你不希望暴露 Kiali UI 通过外部 IP，你也可以使用端口转发直接访问 Kiali UI。

   运行以下命令将 Kiali 服务的端口转发到本地：

   ```
   bash
   
   
   复制代码
   kubectl port-forward -n istio-system svc/kiali 20001:20001
   ```

   然后，在浏览器中访问 `http://localhost:20001` 来查看 Kiali 的 Web UI。

3. **通过浏览器访问 Kiali：**

   如果你通过外部 IP 访问 Kiali，使用以下 URL：

   ```
   arduino
   
   
   复制代码
   http://<external-ip>/kiali
   ```

   你应该能够看到 Kiali 的登录页面或直接进入 UI。

### 请求服务的api

在浏览器或者postman请求服务的api，模拟用户真实访问

```
http://localhost:30001/hello
http://localhost:30001/hello-go-clinet
http://localhost:30002/hello
http://localhost:30002/hello-go-server
http://localhost:30002/hello-java
http://localhost:30003/hello
http://localhost:30003/hello-go-server

```

### 在kiali控制台查看服务拓扑图

http://localhost:20001/ 可以看到服务的调用拓扑图

![image-20250109171154930](/Users/shishupei/Library/Application Support/typora-user-images/image-20250109171154930.png)