#!/bin/bash

function swagger() {
  git clone --depth 1 https://github.com/swagger-api/swagger-ui.git swagger-ui-temp && \
      cp -R swagger-ui-temp/dist swagger-ui && \
      rm -rf swagger-ui-temp
}

"$@"