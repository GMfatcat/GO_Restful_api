<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

[![Static Badge](https://img.shields.io/badge/Python-3.10-green)
](https://www.python.org/)
[![Static Badge](https://img.shields.io/badge/Go-1.19-green)
](https://go.dev/)
[![Static Badge](https://img.shields.io/badge/PostgreSQL-16.1-green)
](https://www.postgresql.org/)

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Architecture & Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Main Page](/repo_img/mainpage.png)

This is the simple fullstack project using **Go,Python,SQL,HTML,CSS and Javascript**.
    
User can query dates to get & save json data, aslo leaflet maps.

Functions of the project :smile:
* Generating fake geo-data in json format --> Python
* Store geo-data in sql table --> Go, PostgreSQL
* Lightweight HTTP Server --> Go
* Interactive Web Page --> HTML,CSS and Javascript

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [![Static Badge](https://img.shields.io/badge/Python-3.7-green)
](https://www.python.org/)
* [![Static Badge](https://img.shields.io/badge/Pygame-2.1.2-green)
](https://www.pygame.org/news)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* Language Version
  ```
  Python >= 3.7
  Go >= 1.12 (using go mod)
  ```
* Package
pq : Go PostgreSQL Driver
  ```
  go get github.com/lib/pq
  ```

### Installation

_Git clone the project or download the whole ZIP file._

1. Git
   ```
   git clone https://github.com/GMfatcat/GO_Restful_api.git
   ```
2. Download the zip file with the bellow instruction
   ```
   1. Click <Code> (green button)
   2. Download ZIP
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Architecture & Usage

![Architecture](/repo_img/architecture.png)

Architecture:
* Frontend (User + Homepage + Result page)
* Backend (HTTP Server + SQL Server + File Server)
* Data Provider (Generate json data)

Data Provider:
* Argparse script for generating data, see arguments info below
   ```
   usage: generate_data.py [-h] [--max_position MAX_POSITION] [--max_num MAX_NUM] [--days DAYS] [--start_date START_DATE]

   options:
  -h, --help            show this help message and exit
  --max_position MAX_POSITION Maximum number of positions to generate.
  --max_num MAX_NUM     Maximum number of objects to generate.
  --days DAYS           Number of days between to generate.
  --start_date START_DATE Start date in the format YYYY-MM-DD.

  Quick example: python generate_data.py --max_position 30 --max_num 30 --days 20 --start_date 2023-01-05
   ```
* Modify coordinates range if necessary
   ```
   edit in generate_data.py line 24 & 25
   ```
* Run the script without any arguments is fine **but those data has already been generated in "./data" folder**
* Save json in "./data"

Backend Usage:
* Configure your PostgreSQL database user / dbname / password in **"./action/connector/Connector.go"**
* Add json data to SQL database
   ```
   // change the *addNewData* to true in add_data.go, default to be false
   go run add_data.go
   ```
* Query dates for sql data
   ```
   // v1 & v2 differences see test_backend.go
   go run test_backend.go
   ```
* Run Http Server - **Must run before ANY Frontend Usage!!!**
   ```
   // After start the server , Open browser, type "http://localhost:9487" in url section.
   go run server.go.
   ```

Frontend Usage ([example video](https://youtu.be/cUN5b4UNa5A)):
* Connect to Homepage
   ```
   Open browser, type "http://localhost:9487" in url section.
   You can change the port in "server.go"
   ```
* Query dates for json data (Date format:XXXX-XX-XX)
* Auto-display json data (if data exists)
* Save json & Generate Leaflet map by button

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

I am trying to combine this project with my another geohash project.  
That project can use geohash to find several groups of what3words.  

- [x] Data Generator
- [x] SQL Server for Data
- [x] Go HTTP Server
- [ ] Backend Functions
    - [x] JSON to SQL Server
    - [x] Query SQL TABLE
    - [x] Middleware Logger
    - [ ] Update / Delete TABLE
    - [ ] TBD
- [ ] Frontend Functions
    - [x] Webpage
    - [x] Query data
    - [x] Save JSON data
    - [x] Interactive Map (Leaflet)
    - [ ] Direct "curl / get" to access data
    - [ ] TBD

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

GMfatcat : [github](https://github.com/GMfatcat)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

I've included all of free resources to kick things off!

* [Mainpage Background Cat Image](https://playgroundai.com/)
* [Webpage Template](https://bootstrapmade.com/append-bootstrap-website-template/)
* [Architecture Image](https://excalidraw.com/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>