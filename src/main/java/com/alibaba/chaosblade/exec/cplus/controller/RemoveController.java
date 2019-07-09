package com.alibaba.chaosblade.exec.cplus.controller;

import com.alibaba.chaosblade.exec.cplus.common.Constants;
import com.alibaba.chaosblade.exec.cplus.common.Response;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.module.MappingBean;
import com.alibaba.chaosblade.exec.cplus.utils.ExecShellUtils;
import com.alibaba.chaosblade.exec.cplus.utils.StringUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author Pengfei Zhou
 */
@RestController
public class RemoveController {
    private static final Logger logger = LoggerFactory.getLogger(RemoveController.class);

    @Value("${script.location}")
    private String strScriptLocation;

    @Value("${script.remove.process.file.name}")
    private String strScriptRemoveProcessFileName;

    @Value("${current.process.name}")
    private String strCurrentProcessName;

    @RequestMapping("/remove")
    public Response remove() {
        logger.info("Start to remove");

        String processName = MappingBean.getInstance().getProcessName();
        if (!StringUtil.isBlank(processName)) {
            Result removeResult = ExecShellUtils.execShell(strScriptLocation + strScriptRemoveProcessFileName, processName);
            if (!removeResult.isSuccess()) {
                logger.info("Fail to destroy, processName: " + processName +
                        "Error code: " + Response.Code.SERVER_ERROR + " error message: " + removeResult.getErr());
            } else {
                MappingBean.getInstance().setSuid("");
                MappingBean.getInstance().setProcessName("");

                logger.info("Succeed to destroy, processName: " + processName);
            }
        }

        logger.info("Start to exit");
        System.exit(0);

        return Response.ofSuccess("Succeed to remove");
    }
}