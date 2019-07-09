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
