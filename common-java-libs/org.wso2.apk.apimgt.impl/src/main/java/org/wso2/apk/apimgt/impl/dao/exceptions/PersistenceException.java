/*
 * Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.apk.apimgt.impl.dao.exceptions;

import org.wso2.apk.apimgt.api.ErrorHandler;
import org.wso2.apk.apimgt.api.ExceptionCodes;

/**
 * This is the custom exception class for API Persistence.
 */
public class PersistenceException extends Exception {

    private ErrorHandler errorHandler;

    /**
     * Get error handler object.
     * @return ErrorHandler
     */
    public ErrorHandler getErrorHandler() {
        return errorHandler;
    }

    public PersistenceException(String msg) {
        super(msg);
        this.errorHandler = ExceptionCodes.INTERNAL_ERROR;
    }

    public PersistenceException(String msg, Throwable e) {
        super(msg, e);
        this.errorHandler = ExceptionCodes.INTERNAL_ERROR;
    }

    public PersistenceException(Throwable throwable) {
        super(throwable);
        this.errorHandler = ExceptionCodes.INTERNAL_ERROR;
    }

    /**
     * This is a default constructor where you can pass error code to error DTO
     * @param message Error message
     * @param code Exception code that need to pass to the error DTO
     */
    public PersistenceException(String message, ErrorHandler code) {
        super(message);
        this.errorHandler = code;
    }

    /**
     * This is a default constructor where you can pass error code to error DTO
     * @param code Exception code that need to pass to the error DTO
     */
    public PersistenceException(ErrorHandler code) {
        super(code.getErrorCode() + ":" + code.getErrorMessage() + "::" + code.getErrorDescription());
        this.errorHandler = code;
    }

    /**
     * This is a default constructor where you can pass error code to error DTO
     * @param message Error message
     * @param cause throwable object.
     * @param code Exception code that need to pass to the error DTO
     */
    public PersistenceException(String message, Throwable cause, ErrorHandler code) {
        super(message, cause);
        this.errorHandler = code;
    }
}
