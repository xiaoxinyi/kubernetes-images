# FROM haskell
# WORKDIR /opt/hello_name
# RUN cabal user-config init
# RUN echo "\
# repository mirrors.ustc.edu.cn\n\
#     url: https://mirrors.ustc.edu.cn/hackage/\n\
#     secure: True\n\
# " > ~/.cabal/config
# RUN cabal update
# COPY ./haskell_example.cabal /opt/hello_name/haskell_example.cabal
# RUN cabal build --only-dependencies -j4
# COPY . /opt/hello_name
# RUN cabal install
# CMD ["haskell_example"]

FROM golang:buster

WORKDIR /workspace
COPY main.go /workspace/
RUN go build main.go
CMD ["main"]