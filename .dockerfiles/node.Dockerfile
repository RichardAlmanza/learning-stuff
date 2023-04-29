# syntax=docker/dockerfile-upstream:master-labs

ARG DOCKER_IMAGE_TAG=16.18-alpine3.16
ARG GIT_VERSION=2.36.3
ARG ZSH_VERSION=5.8.1

FROM node:${DOCKER_IMAGE_TAG}

ARG USERNAME=node
ARG OMZ_PATH=/etc/omz/.oh-my-zsh
ARG GIT_VERSION
ARG ZSH_VERSION
ARG NPM_GLOBAL=/usr/local/share/npm-global

# Add NPM global to PATH.
ENV PATH=${NPM_GLOBAL}/bin:${PATH}

ADD https://github.com/ohmyzsh/ohmyzsh.git ${OMZ_PATH}

SHELL [ "/bin/sh", "-o", "pipefail", "-c" ]

RUN \
    # Set zshrc config
    cp ${OMZ_PATH}/templates/zshrc.zsh-template /root/.zshrc \
    && sed -i 's/robbyrussell/ys/g; s/# HIST_STAMPS/HIST_STAMPS/g' /root/.zshrc \
    && cp /root/.zshrc /home/${USERNAME}/.zshrc \
    && chown ${USERNAME}:${USERNAME} /home/${USERNAME}/.zshrc \
    # Create links for Oh My Zsh
    && ln -s ${OMZ_PATH} /root/.oh-my-zsh \
    && su ${USERNAME} -c "ln -s ${OMZ_PATH} /home/${USERNAME}/.oh-my-zsh" \
    # Configure global npm install location, use group to adapt to UID/GID changes
    && if ! grep -e "^npm:" /etc/group > /dev/null 2>&1; then addgroup -S npm; fi \
    && addgroup ${USERNAME} npm \
    && umask 0002 \
    && mkdir -p ${NPM_GLOBAL} \
    && mkdir -p /usr/local/etc/ \
    && touch /usr/local/etc/npmrc \
    && chown ${USERNAME}:npm ${NPM_GLOBAL} /usr/local/etc/npmrc \
    && chmod g+s ${NPM_GLOBAL} \
    && npm config -g set prefix ${NPM_GLOBAL} \
    && su ${USERNAME} -c "npm config -g set prefix ${NPM_GLOBAL}"

RUN \
    # Install eslint
    su ${USERNAME} -c "umask 0002 && npm install -g eslint@8.26.x " \
    && npm cache clean --force > /dev/null 2>&1 \
    # Install packages
    && apk add --no-cache \
    git=~${GIT_VERSION} \
    zsh=~${ZSH_VERSION} \
    # Set default Shell to zsh
    && sed -i 's/\/bin\/\(ash\|sh\|bash\)/\/bin\/zsh/g' /etc/passwd
