FROM bitnami/kubectl
ARG APK_USER=wso2
ARG APK_USER_ID=10001
ARG APK_USER_GROUP=wso2
ARG APK_USER_GROUP_ID=10001
ARG APK_USER_HOME=/home/${APK_USER}

USER 0

RUN apt-get update \
 && apt-get install -y bash bc\
 && addgroup --system -gid ${APK_USER_GROUP_ID} ${APK_USER_GROUP} \
 && adduser --system -u ${APK_USER_ID} -h ${APK_USER_HOME} -G ${APK_USER_GROUP} ${APK_USER} \
 && mkdir -p ${APK_USER_HOME} \
 && chown -R ${APK_USER_ID}:${APK_USER_GROUP_ID} ${APK_USER_HOME} \
 && chgrp -R 0 ${APK_USER_HOME} \
 && chmod -R g=u ${APK_USER_HOME}

ENV PATH $PATH:/usr/local/go/bin

USER 10001

WORKDIR ${APK_USER_HOME}
COPY scripts/ /scripts/

ENTRYPOINT ["bash", "/scripts/entrypoint.sh"]
