# 🤝 Recoverly - Mental Health & Recovery Platform for India 🇮🇳

## 🚀 Overview

Recoverly is a mobile and web platform designed to provide accessible and effective mental health support and recovery resources specifically tailored for the Indian context. 🧘‍♀️ It aims to address various aspects of mental well-being through support, community engagement, and personalized recovery plans.

![banner-modified](https://github.com/user-attachments/assets/59456af3-7e22-4e26-9f47-ed78ffe27f1d)


## ✨ Key Features

* **📞 Support and Rehabilitation:** Access an emergency hotline for immediate assistance.
* **👥 Youth Engagement:** Connect with peer support groups for shared experiences and guidance.
* **🛠️ Community Building:** Utilize a workshop kit generator to facilitate community-driven support.
* **📝 Treatment and Recovery:** Benefit from personalized recovery plans and dashboards for progress tracking.
* **📅 Event Coordination:** Streamlined event participation with QR code check-ins and an event finder.

## 💡 Key Impact & Innovation

* **Saving Lives and Reducing Costs:**
    * Helps prevent relapses.
    * Saves lives.
    * Eases economic burdens through scalable, low-cost recovery solutions.
* **Personalized Recovery Plans:**
    * Leverages AI to deliver tailored recovery plans.
    * Provides actionable goals.
    * Offers real-time progress tracking.
    * Ensures every user gets a unique and effective path to sobriety.

## 🛠️ Tech Stack

* **Frontend:** Svelte, CSS, Tailwind CSS
* **Backend:** Go
* **Database:** PostgreSQL
* **Deployment:** Raspberry Pi
* **Server Management:** NGINX
* **CI/CD:** Git
* **Bundling:** Capacitor JS

## ⚙️ Setup Instructions

1.  **Prerequisites:**
    * Node.js and npm installed.
    * Go installed and configured.
    * PostgreSQL installed and running.
    * Capacitor CLI installed (`npm install -g @capacitor/cli`).
2.  **Clone the Repository:**
    ```bash
    git clone https://github.com/hxri-nxrxyxn/recoverly-sv
    cd Recoverly
    ```
3.  **Backend Setup (Go):**
    * Navigate to the `api` directory.
    * Configure your PostgreSQL connection details.
    * Run `go mod tidy` to install dependencies.
    * Run `go build` to compile the backend.
    * Run the backend: `./main`.
      
4.  **Database Setup (PostgreSQL):**
    * Create a PostgreSQL database for Recoverly.
    * Run the necessary SQL scripts to create the tables.
      
5.  **Frontend Setup (Svelte):**
    * Navigate to the `src` directory.
    * Run `npm install` to install dependencies.
    * Run `npm run dev` to start the development server.
      
6.  **Capacitor Setup (Android):**
    * In the `frontend` directory, run `npx cap init`.
    * Add Android platform: `npx cap add android`.
    * Build the Svelte app: `npm run build`.
    * Copy the built files to the Capacitor web directory: `npx cap copy`.
    * Open Android Studio: `npx cap open android`.
      
7.  **Deployment:**
    * Follow your Raspberry Pi setup instructions for deploying the backend and frontend.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to suggest improvements or report bugs. 🐛

## 📜 License

This project is licensed under the MIT License. 📝

## 📧 Contact

For any questions or inquiries, please contact [me](hari@laddu.cc). 📬
