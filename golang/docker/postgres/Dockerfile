FROM postgres:10.17
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
COPY init.sql /docker-entrypoint-initdb.d/init.sql
RUN chmod 0755 /docker-entrypoint-initdb.d/init.sql
