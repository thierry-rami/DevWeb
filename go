#!/bin/bash
 read -p "Saisissez l'adresse mail utilisée avec votre compte Github : " email
 read -p "Saisissez le nom du compte github : " gaccount
 read -p "Saisissez le token que vous as fournis Github : " token
 read -p "Saisissez le mot de passe (gmail) :" mdpgmail

rm .ssh/id*
ssh-keygen -t rsa -b 4096 -f "$HOME"/.ssh/id_ed25519 -C "$email" -N ""

git config --global user.email "$email"
git config --global user.name "$gaccount"
git config --global init.defaultBranch main
dt=$(date '+%Y-%m-%d')

curl -X POST -H "Authorization: token $token" \
             -d '{"title" : "SSH_Key '$dt'  Plateforme", "key" : "'"$(cat "$HOME"/.ssh/id_ed25519.pub)"'"}' \
             https://api.github.com/user/keys

sed -i '/export Git_Token/d' .bashrc
echo 'export Git_Token="'$token'"' >> .bashrc
source .bashrc

sudo sed -i '/from/d' /etc/msmtprc
sudo sed -i '/user/d' /etc/msmtprc
sudo sed -i '/password/d' /etc/msmtprc
sudo sed -i '/account default/d' /etc/msmtprc
sudo chmod 666 /etc/msmtprc
sudo echo "from $email"  >> /etc/msmtprc
sudo echo "user $email"  >> /etc/msmtprc
sudo echo "password $mdpgmail"  >> /etc/msmtprc
sudo echo ""  >> /etc/msmtprc
sudo echo "account default : gmail" >> /etc/msmtprc

# test du mail
 printf "Subject:msmtp gmail\nL\'envoi de mail Fonctionne !" | msmtp $email
