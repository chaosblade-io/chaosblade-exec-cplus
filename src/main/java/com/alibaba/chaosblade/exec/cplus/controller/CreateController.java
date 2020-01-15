/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.alibaba.chaosblade.exec.cplus.controller;

import com.alibaba.chaosblade.exec.cplus.common.Constants;
import com.alibaba.chaosblade.exec.cplus.module.MappingBean;
import com.alibaba.chaosblade.exec.cplus.module.RequestParamsBean;
import com.alibaba.chaosblade.exec.cplus.service.CheckDuplicateService;
import com.alibaba.chaosblade.exec.cplus.service.CheckStatusService;
import com.alibaba.chaosblade.exec.cplus.service.ModelFactory;
import com.alibaba.chaosblade.exec.cplus.service.IModelService;
import com.alibaba.chaosblade.exec.cplus.common.Response;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.utils.StringUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.*;

/**
 * @author Pengfei Zhou
 */
@RestController
public class CreateController {

    private static final Logger logger = LoggerFactory.getLogger(CreateController.class);

    @Autowired
    private ModelFactory modelFactory;
    @Autowired
    private CheckDuplicateService checkDuplicateService;

    @RequestMapping("/create")
    @ResponseBody
    public Response create(final RequestParamsBean requestParamsBean)  {
        logger.info("Start to create, requestParamsBean: "+ requestParamsBean.toString());

        // check necessary arguments
        String suid = requestParamsBean.getSuid();
        if (StringUtil.isBlank(suid)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less experiment argument");
        }
        String target = requestParamsBean.getTarget();
        if (StringUtil.isBlank(target)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less target argument");
        }
        String actionArg = requestParamsBean.getAction();
        if (StringUtil.isBlank(actionArg)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less action argument");
        }
        String breakLine = requestParamsBean.getBreakLine();
        if (StringUtil.isBlank(breakLine)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less breakLine argument");
        }
        String fileLocateAndName = requestParamsBean.getFileLocateAndName();
        if (StringUtil.isBlank(fileLocateAndName)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less fileLocateAndName argument");
        }
        String forkMode = requestParamsBean.getForkMode();
        if (StringUtil.isBlank(forkMode)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less forkMode argument");
        }
        String processName = requestParamsBean.getProcessName();
        if (StringUtil.isBlank(processName)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "less processName argument");
        }

        // check the command supported or not
        if (!target.equals(Constants.TARGET_NAME)) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, "the target not supported");
        }

        String libLoad = requestParamsBean.getLibLoad();
        if (!StringUtil.isBlank(libLoad)) {
            requestParamsBean.setLibLoad(Constants.SET_ENV_LD_LIBRARY_PATH + libLoad);
        } else {
            requestParamsBean.setLibLoad("");
        }

        boolean isDuplicate = checkDuplicateService.isExistProcessByProcessName(processName);
        if(isDuplicate){
            return Response.ofFailure(Response.Code.DUPLICATE_INJECTION, "the experiment exists");
        }

        final IModelService iModelService = modelFactory.createModelService(actionArg);
        if(iModelService == null){
            return Response.ofFailure(Response.Code.NOT_FOUND, "the action not supported");
        }

        Result validateResult = iModelService.validateFlag(requestParamsBean);
        if (!validateResult.isSuccess()) {
            return Response.ofFailure(Response.Code.ILLEGAL_PARAMETER, validateResult.getErr());
        }

        MappingBean.getInstance().setSuid(suid);
        MappingBean.getInstance().setProcessName(processName);

        // begin to inject
        final Response response = new Response();
        ExecutorService executorService = new ThreadPoolExecutor(Constants.CORE_POOL_SIZE, Constants.MAXIMUM_POOL_SIZE,
                Constants.KEEP_ALIVE_TIME, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>());
        executorService.execute(new Runnable() {
                                    @Override
                                    public void run() {
                                        Result injectionResult = iModelService.handleInjection(requestParamsBean);
                                        if (!injectionResult.isSuccess()) {

                                            MappingBean.getInstance().setSuid("");
                                            MappingBean.getInstance().setProcessName("");

                                            logger.info("Fail to create, requestParamsBean: " + requestParamsBean.toString() +
                                                    "Error code: " + Response.Code.SERVER_ERROR + " error message: " + injectionResult.getErr());

                                            response.setCode(Response.Code.SERVER_ERROR.getCode());
                                            response.setError(injectionResult.getErr());
                                            response.setSuccess(false);

                                        } else {
                                            logger.info("Succeed to create, requestParamsBean: " + requestParamsBean.toString());
                                        }
                                    }
                                });

        executorService.shutdown();

        try {
            Thread.sleep(Constants.SLEEP_MILLISECONDS);
            if (response.getError() != null) {
                return response;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return Response.ofSuccess(requestParamsBean.toString());
    }
}
