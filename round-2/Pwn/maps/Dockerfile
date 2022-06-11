FROM ubuntu:20.04

ENV USER anonuser

WORKDIR /home/$USER

RUN useradd $USER
RUN apt-get update
RUN apt-get install -y xinetd
RUN apt-get update
RUN apt-get install -y build-essential
RUN apt-get update
RUN apt-get install -y gcc-multilib

COPY maps.xinetd /etc/xinetd.d/maps
COPY src/maps.c /home/$USER/
COPY src/Makefile /home/$USER/Makefile
COPY flag.txt /home/$USER/flag.txt

RUN make -C /home/$USER/

RUN chown -R root:$USER /home/$USER
RUN chmod -R 550 /home/$USER
RUN chmod +r /home/$USER/flag.txt

EXPOSE 17712

CMD service xinetd start && sleep 2 && tail -f /var/log/xinetdlog
