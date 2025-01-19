<script>
    import { Link } from "svelte-routing";
    import { Storage } from "@capacitor/storage";
    const urlParams = new URLSearchParams(window.location.search);
    const duration = urlParams.get("duration"); // Get the value of 'category'
    import Nav from "./Nav.svelte";
    import { GoogleGenerativeAI } from "@google/generative-ai";
    const API_KEY = "AIzaSyCjT3qZw8I6SYEqEH1x_681_4czeQB8OIw";
    let mytip = ""; // Get the value of 'category'
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
        `generate a short term plan for a person to recover from an addiction., only use plain text and limited characters, (something like prepare to quit smoking in 60 days -> high chances of success, each day reduce one chance)`,
    );
</script>

<main>
    <Nav />
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="gen__box">
        <h3>Short Term Goal</h3>
        <p>{mytip}</p>
    </div>
    <div class="btn__box">
        <button class="outlined" onclick={generateAI}>REGENERATE</button>
        <br />
        <Link to={`/dashboard`}>
            <button class="solid">SAVE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
