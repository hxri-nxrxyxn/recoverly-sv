const baseUrl = "https://hardly-genuine-yeti.ngrok-free.app";

async function get(endpoint) {
    try {
        const response = await fetch(`${baseUrl}${endpoint}`);
        if (!response.ok) {
            throw new Error(`Network response was not ok (${response.status})`);
        }
        return await response.json();
    } catch (error) {
        alert("Error fetching data:", error);
        throw error; // Re-throw the error for proper handling
    }
}

async function post(endpoint, data) {
    try {
        const response = await fetch(`${baseUrl}${endpoint}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
        if (!response.ok) {
            console.log(`Network response was not ok (${response})`);
        }
        return await response.json();
    } catch (error) {
        alert("Error creating data:", error);
        throw error;
    }
}

async function put(endpoint, data) {
    try {
        const response = await fetch(`${baseUrl}${endpoint}`, {
            method: "PUT",
            // headers: {
            // 'Content-Type': 'application/json',
            // },
            body: JSON.stringify(data),
        });
        if (!response.ok) {
            throw new Error(`Network response was not ok (${response.status})`);
        }
        return await response.json();
    } catch (error) {
        alert("Error updating data:", error);
        throw error;
    }
}

async function del(endpoint) {
    try {
        const response = await fetch(`${baseUrl}${endpoint}`, {
            method: "DELETE",
        });
        if (!response.ok) {
            throw new Error(`Network response was not ok (${response.status})`);
        }
        return response.status === 204; // Check for successful deletion (no content)
    } catch (error) {
        alert("Error deleting data:", error);
        throw error;
    }
}

// Example Usage
async function fetchData() {
    try {
        const data = await get("/users");
        alert(data);

        const newData = { name: "New User", email: "newuser@example.com" };
        const createdData = await post("/users", newData);
        alert(createdData);

        const updatedData = { id: createdData.id, name: "Updated User" };
        const updatedResult = await put(`/users/${updatedData.id}`, updatedData);
        alert(updatedResult);

    } catch (error) {
        alert("An error occurred:", error);
    }
}

export { get, post, put, del };