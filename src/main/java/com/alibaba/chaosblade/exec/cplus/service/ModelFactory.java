package com.alibaba.chaosblade.exec.cplus.service;

import com.alibaba.chaosblade.exec.cplus.common.Constants;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

/**
 * @author Pengfei Zhou
 */
@Component
public class ModelFactory {

    @Autowired
    private DelayModelService delayModelService;
    @Autowired
    private ReturnErrorModelService returnErrorModelService;
    @Autowired
    private VariableModifyModelService variableModifyModelService;

    public IModelService createModelService(String actionName) {
        if(Constants.DELAY_ACTION_NAME.equals(actionName)){
            return delayModelService;
        } else if(Constants.RETURN_ERROR_DATA_ACTION_NAME.equals(actionName)){
            return returnErrorModelService;
        } else if(Constants.VARIABLE_MODIFY_ACTION_NAME.equals(actionName)){
            return variableModifyModelService;
        }else {
            return null;
        }
    }
}
