FROM minio/minio:RELEASE.2024-06-13T22-53-53Z-cpuv1

COPY create-buckets.sh /usr/local/bin/create-buckets.sh
RUN chmod +x /usr/local/bin/create-buckets.sh

ENTRYPOINT ["/usr/local/bin/create-buckets.sh"]
