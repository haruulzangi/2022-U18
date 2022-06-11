FROM ubuntu:18.04

ENV USER anonuser

WORKDIR /home/$USER

RUN useradd $USER
RUN apt-get update
RUN apt-get install -y xinetd
RUN apt-get update
RUN apt-get install -y build-essential
RUN apt-get update
RUN apt-get install -y gcc-multilib

COPY warmup.xinetd /etc/xinetd.d/warmup
COPY src/warmup.c /home/$USER/
COPY src/Makefile /home/$USER/Makefile
COPY flag.txt /home/$USER/flag.txt

RUN make -C /home/$USER/

RUN chown -R root:$USER /home/$USER
RUN chmod -R 550 /home/$USER
RUN chmod +r /home/$USER/flag.txt

EXPOSE 17711

CMD service xinetd start && sleep 2 && tail -f /var/log/xinetdlog
