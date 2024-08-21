# Attendance Tracker

A simple and efficient attendance tracking application built with Go, Redis, and a frontend interface. The app allows users to manage subjects, update attendance records, and ensure compliance with minimum attendance requirements.

## Features

- **User Registration:** Securely register users.
- **Subject Management:** Add, delete, and manage subjects like tasks in a to-do app.
- **Attendance Tracking:** Increment or decrement present and working days for each subject.
- **Real-Time Updates:** Dynamically update attendance data and save changes with a single click.
- **Minimum Attendance Alert:** Warn users if attendance drops below 75%.

## Tech Stack

- **Backend:** Go, Redis
- **Frontend:** HTML, CSS, JavaScript
- **Database:** Hosted on Heroku
- **Deployment:** Vercel (Frontend), Heroku (Backend)

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/attendance-tracker.git
    cd attendance-tracker
    ```

2. **Set up Redis and Go dependencies:**
    ```bash
    go mod download
    ```

3. **Run the application:**
    ```bash
    go run cmd/server/main.go
    ```

4. **Access the app:**
   Open `http://localhost:8080` in your browser.

   
#### Note :  currently not deployed due to some bugs, will be soon back 😉

## Deployment

- **Frontend:** Deployed on Vercel.
- **Backend:** Hosted on Heroku.

## Usage

- **Add Subject:** Use the form to add new subjects.
- **Manage Attendance:** Increment or decrement present and working days, and save updates.
- **Delete Subject:** Remove subjects as needed.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.
