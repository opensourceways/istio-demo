package org.shishupei.javaistioserver;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class HelloController {
    @Autowired
    private RestTemplate restTemplate;

    @GetMapping("/hello")
    public String hello() {
        return "Hello, World JAVA!";
    }

    @GetMapping("/hello-go-server")
    public String callGoIstioHello() {
        // 通过 Kubernetes 服务名称访问 go-istio-server 的 hello API
        String goIstioUrl = "http://go-istio-server.default.svc.cluster.local/hello";
        return restTemplate.getForObject(goIstioUrl, String.class);
    }
}

