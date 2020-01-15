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

package com.alibaba.chaosblade.exec.cplus.common;

/**
 * @author Pengfei Zhou
 */
public class Result {
    private boolean success;
    private String err;
    private String comment;

    public Result(boolean success, String err, String comment) {
        this.success = success;
        this.err = err;
        this.comment = comment;
    }

    public static Result fail(String err) {
        return new Result(false, err, null);
    }

    public static Result success() {
        return new Result(true, null, null);
    }

    public static Result success(String comment) {
        return new Result(true, null, comment);
    }

    public boolean isSuccess() {
        return success;
    }

    public void setSuccess(boolean success) {
        this.success = success;
    }

    public String getErr() {
        return err;
    }

    public void setErr(String err) {
        this.err = err;
    }

    public String getComment() {
        return comment;
    }

    public void setComment(String comment) {
        this.comment = comment;
    }
}
