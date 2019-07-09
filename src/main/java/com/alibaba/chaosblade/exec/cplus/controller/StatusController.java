package com.alibaba.chaosblade.exec.cplus.controller;

import com.alibaba.chaosblade.exec.cplus.common.Response;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author Pengfei Zhou
 */
@RestController
public class StatusController {

    @RequestMapping("/status")
    public Response status(){
        return Response.ofSuccess("succeed to start the program");
    }
}
