FROM alpine:latest
ADD circleapi /
ADD config.yaml /config.yaml
ENV CIRCLE_TMPDIR="/tempdata/"  CIRCLE_ENDPOINT="http://oss-cn-qingdao.aliyuncs.com" CIRCLE_ACCESSKEY="y1212" \
CIRCLE_SECRET="23232323223" CIRCLE_BUCKET="antzbdev" CIRCLE_DBUSER="postgres" CIRCLE_DBPWD="Antzb"\
CIRCLE_DBHOST="12.3.4.5" CIRCLE_DBNAME="antzb"
VOLUME ["/tempdata"]
CMD ["/circleapi"]
