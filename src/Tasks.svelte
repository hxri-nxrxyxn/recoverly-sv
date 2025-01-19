<script>
    import { Link } from "svelte-routing";
    import { get } from "./fetch";
    import { GoogleGenerativeAI } from "@google/generative-ai";
    const API_KEY = "AIzaSyCjT3qZw8I6SYEqEH1x_681_4czeQB8OIw";
    import Nav from "./Nav.svelte";

    const urlParams = new URLSearchParams(window.location.search);
    let mytip = urlParams.get("tip"); // Get the value of 'category'
    const generateAI = async (prompt) => {
        const genAI = new GoogleGenerativeAI(API_KEY);
        const model = genAI.getGenerativeModel({
            model: "gemini-1.5-flash",
        });
        try {
            const result = await model.generateContent(prompt);
            mytip = result.response.text();
            console.log(mytip);
        } catch (error) {}
    };

    generateAI(
        `dont't even have to include any backticks just return the data in a array format, as a series of tasks (max 3) with list of having whaat to do`,
    );
</script>

<main>
    <Nav />
    <h1>Generate Effective <br /> <span>Goals</span></h1>

    <div class="tasks">
        <div class="task">
            <h3>Week 1</h3>
            <p>Get started by</p>
            <p class="task--sub">task 1</p>
            <p class="task--sub">task 2</p>
        </div>
    </div>
    <div class="btn__box">
        <Link to="/dashboard">
            <button class="solid">CONTINUE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
