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

package org.wso2.apk.apimgt.api.model;

import java.util.ArrayList;
import java.util.List;

public class CommentList {

    private Integer count = null;
    private List<Comment> list = new ArrayList<Comment>();
    private Pagination pagination = null;

    public Integer getCount() {return count; }

    public void setCount(Integer count) {this.count = count; }

    public List<Comment> getList() {return list; }

    public void setList(List<Comment> list) {this.list = list; }

    public Pagination getPagination() {return pagination; }

    public void setPagination(Pagination pagination) {this.pagination = pagination; }
}
