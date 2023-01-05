FROM google/cloud-sdk:412.0.0-alpine

# Install go
RUN apk add wget openjdk11 && \
    wget -c https://go.dev/dl/go1.15.10.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.15.10.linux-amd64.tar.gz

# Set go globally
ENV PATH="$PATH:/usr/local/go/bin"

# Enbale quite mode
ENV CLOUDSDK_CORE_DISABLE_PROMPTS=1

# Install and use python2.7 as default since that is required by the dev_appserver.py script.
# For reference see https://cloud.google.com/appengine/docs/standard/tools/using-local-server?tab=go#running_the_local_development_server
# Install pip and install grpcio for python2 otherwise it error will occur when running the dev_appserver.py script.
RUN apk -q add python2

# APPLCATION_ID is required for the datastore emulator and needs to be set.
# When openeing the admin panel it is on the left side of the banner.
ENV APPLICATION_ID=dev~None

RUN echo -n "beta app-engine-go cloud-datastore-emulator app-engine-python-extras" > /tmp/additional_components
RUN /google-cloud-sdk/install.sh --bash-completion=false --path-update=true --usage-reporting=false \
	--additional-components `cat /tmp/additional_components` && rm -rf /google-cloud-sdk/.install/.backup
RUN git config --system credential.'https://source.developers.google.com'.helper gcloud.sh
