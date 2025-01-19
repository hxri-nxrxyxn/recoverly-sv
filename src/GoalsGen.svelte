<script>
    import { Link } from "svelte-routing";
    import { Storage } from "@capacitor/storage";
    const urlParams = new URLSearchParams(window.location.search);
    const duration = urlParams.get("duration"); // Get the value of 'category'
    import { GoogleGenerativeAI } from "@google/generative-ai";
    const API_KEY = "AIzaSyCjT3qZw8I6SYEqEH1x_681_4czeQB8OIw";
    import Nav from "./Nav.svelte";

    let mytip = $state("");
    let hasToken = null;
    const checkUser = async () => {
        try {
            const value = await Storage.get({ key: "id" });
            const id = value.value;
            if (!id) {
                hasToken = false;
                location.href = "/login";
            }
            const response = await fetch(
                `http://192.168.157.224:8080/api/v1/users/${id}`,
                {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                },
            );

            if (response.ok) {
                const data = await response.json();
                console.log(data.data);
                const generateAI = async (prompt) => {
                    const genAI = new GoogleGenerativeAI(API_KEY);
                    const model = genAI.getGenerativeModel({
                        model: "gemini-1.5-flash",
                    });
                    try {
                        const result = await model.generateContent(prompt);
                        mytip = result.response.text();
                    } catch (error) {}
                };

                generateAI(
                    `this is the data provided by a patient of ours, he is trying to recover his life, please give him a short term plan for recovery, like reduciinig whatever he is using day by and all. here are his details ${data.data.name} and his preferrable duration ${duration}`,
                );
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };

    checkUser();
</script>

<main>
    <Nav />
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="gen__box">
        <h3>Short Term Goal</h3>
        <p>{mytip}</p>
    </div>
    <div class="btn__box">
        <button class="outlined" onclick={checkUser}>REGENERATE</button>
        <br />
        <Link to={`/tasks?tip=${mytip}`}>
            <button class="solid">SAVE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
