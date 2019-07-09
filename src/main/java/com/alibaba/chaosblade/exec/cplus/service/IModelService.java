package com.alibaba.chaosblade.exec.cplus.service;

import com.alibaba.chaosblade.exec.cplus.module.RequestParamsBean;
import com.alibaba.chaosblade.exec.cplus.common.Result;

/**
 * @author Pengfei Zhou
 */
public interface IModelService {

    //validate flag
    Result validateFlag(RequestParamsBean requestParamsBean);

    //handle injection
    Result handleInjection(RequestParamsBean requestParamsBean);
}
