# Webhook

> This project consists of an endpoint to consume webhooks of Pull Request from github and send notifications in specific chats.
<p align="center">
  <img src="/assets/example.gif" height="400">
</p>

### Platforms already integrated
<img height="70" src="https://logodownload.org/wp-content/uploads/2017/11/discord-logo-4-1.png" alt="discord"/></code>
<img height="70" src="https://user-images.githubusercontent.com/5147537/54070671-0a173780-4263-11e9-8946-09ac0e37d8c6.png" alt="slack"/>
<img height="70" src="http://fonts.gstatic.com/s/i/productlogos/chat_round_2020q4/v1/web-96dp/logo_chat_round_2020q4_color_2x_web_96dp.png" alt="google_chat"/>

----
## Architecture
<p align="center">
  <img src="docs/images/web-scraping.jpg" height="440">
</p>

## Built With
<img height="50" src="https://gofiber.io/assets/images/logo.svg" alt="fiber"/>
<img height="50" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/MongoDB_Logo.svg/2560px-MongoDB_Logo.svg.png" alt="mongoDB"/>


### Pre-requisites
* go
  ```sh
  sudo snap install go --classic
  ```
* mongoDB

### Installation
1. Clone the repository
   ```sh
   git clone https://github.com/victorvcruz/social_network_project.git
   ```
2. Install go packages
   ```sh
   go mod download
   ```

## Usage
> Add your url in Repository->Setting->Webhook
<p align="center">
  <img src="/assets/insert_example.gif" height="400">
</p>

* Insert your platform chat and your environment variables in `.env`

### To start execution
* run
   ```sh
   go run main.go
   ```

