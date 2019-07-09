package com.alibaba.chaosblade.exec.cplus.controller;

import com.alibaba.chaosblade.exec.cplus.common.Constants;
import com.alibaba.chaosblade.exec.cplus.common.Response;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.module.MappingBean;
import com.alibaba.chaosblade.exec.cplus.module.DestroyProcessBean;
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
public class DestroyController {
    private static final Logger logger = LoggerFactory.getLogger(DestroyController.class);

    @Value("${script.location}")
    private String strScriptLocation;

    @Value("${script.remove.process.file.name}")
    private String strScriptRemoveProcessFileName;

    @RequestMapping("/destroy")
    public Response destroy(DestroyProcessBean destroyProcessBean)  {
        logger.info("Start to destroy, removeProcessBean: "+ destroyProcessBean.toString());

        // check necessary arguments
        String suid = destroyProcessBean.getSuid();

        if (StringUtil.isBlank(suid)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less experiment argument");
        }

        if (!suid.equals(MappingBean.getInstance().getSuid())){
            return Response.ofFailure(Response.Code.NOT_FOUND, "the experiment not found");
        }

        String processName = MappingBean.getInstance().getProcessName();
        if (StringUtil.isBlank(processName)) {
            return Response.ofFailure(Response.Code.SERVER_ERROR, "the process name mapping with " + suid + " not found");
        }

        Result removeResult = ExecShellUtils.execShell(strScriptLocation + strScriptRemoveProcessFileName, processName);
        if (!removeResult.isSuccess()) {
            logger.info("Fail to destroy, removeProcessBean: " + destroyProcessBean.toString() +
                    "Error code: " + Response.Code.SERVER_ERROR + " error message: " + removeResult.getErr());
            return Response.ofFailure(Response.Code.SERVER_ERROR, removeResult.getErr());
        } else {
            MappingBean.getInstance().setSuid("");
            MappingBean.getInstance().setProcessName("");

            logger.info("Succeed to destroy, removeProcessBean: " + destroyProcessBean.toString());
            return Response.ofSuccess(removeResult.toString());
        }
    }
}
