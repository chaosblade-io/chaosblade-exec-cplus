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

package com.alibaba.chaosblade.exec.cplus.module;

/**
 * @author Pengfei Zhou
 */
public class RequestParamsBean {
    private String suid;

    private String target;

    private String action;

    private String breakLine;

    private String returnValue;

    private String delayDuration;

    private String varaibleName;

    private String varaibleValue;

    private String initParams;

    private String fileLocateAndName;

    private String forkMode;

    private String libLoad;

    private String processName;

    public String getSuid() {
        return suid;
    }

    public void setSuid(String suid) {
        this.suid = suid;
    }

    public String getTarget() {
        return target;
    }

    public void setTarget(String target) {
        this.target = target;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }

    public String getBreakLine() {
        return breakLine;
    }

    public void setBreakLine(String breakLine) {
        this.breakLine = breakLine;
    }

    public String getReturnValue() {
        return returnValue;
    }

    public void setReturnValue(String returnValue) {
        this.returnValue = returnValue;
    }

    public String getDelayDuration() {
        return delayDuration;
    }

    public void setDelayDuration(String delayDuration) {
        this.delayDuration = delayDuration;
    }

    public String getVaraibleName() {
        return varaibleName;
    }

    public void setVaraibleName(String varaibleName) {
        this.varaibleName = varaibleName;
    }

    public String getVaraibleValue() {
        return varaibleValue;
    }

    public void setVaraibleValue(String varaibleValue) {
        this.varaibleValue = varaibleValue;
    }

    public String getInitParams() {
        return initParams;
    }

    public void setInitParams(String initParams) {
        this.initParams = initParams;
    }

    public String getFileLocateAndName() {
        return fileLocateAndName;
    }

    public void setFileLocateAndName(String fileLocateAndName) {
        this.fileLocateAndName = fileLocateAndName;
    }

    public String getForkMode() {
        return forkMode;
    }

    public void setForkMode(String forkMode) {
        this.forkMode = forkMode;
    }

    public String getLibLoad() {
        return libLoad;
    }

    public void setLibLoad(String libLoad) {
        this.libLoad = libLoad;
    }

    public String getProcessName() {
        return processName;
    }

    public void setProcessName(String processName) {
        this.processName = processName;
    }

    @Override
    public String toString() {
        return "RequestParamsBean{" +
                "suid='" + suid + '\'' +
                ", target='" + target + '\'' +
                ", action='" + action + '\'' +
                ", breakLine='" + breakLine + '\'' +
                ", returnValue='" + returnValue + '\'' +
                ", delayDuration='" + delayDuration + '\'' +
                ", varaibleName='" + varaibleName + '\'' +
                ", varaibleValue='" + varaibleValue + '\'' +
                ", initParams='" + initParams + '\'' +
                ", fileLocateAndName='" + fileLocateAndName + '\'' +
                ", forkMode='" + forkMode + '\'' +
                ", libLoad='" + libLoad + '\'' +
                ", processName='" + processName + '\'' +
                '}';
    }
}
