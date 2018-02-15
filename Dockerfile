FROM python:2


WORKDIR /usr/src/
RUN git clone https://github.com/genesixx/coalibot app 
WORKDIR /usr/src/app
RUN pip install --no-cache-dir -r requirements.txt
ENV SLACK_API_TOKEN="xoxb-"
ENV INTRA_CLIENT_ID=""
ENV INTRA_SECRET=""
COPY . .

CMD [ "python", "./coalibot.py" ]
