FROM scratch
EXPOSE 8000
ADD go-msgstore /
CMD ["/go-msgstore"]

