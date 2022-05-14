FROM busybox

COPY nfs-provisioner /

ENTRYPOINT ["/nfs-provisioner"]