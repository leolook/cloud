FROM scratch
MAINTAINER hwt
LABEL cloud 1.0
COPY /main /bin
CMD ["/bin"]