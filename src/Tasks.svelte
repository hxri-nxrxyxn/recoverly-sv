<script>
    import { Link } from "svelte-routing";
    import { get } from "./fetch";
    const doPost = async () => {
        const res = await get("/");
        console.log(res);
    };
    doPost();
</script>

<main>
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>Goal duration</label>
            <input type="text" placeholder="Weeks" required />
        </div>
    </div>

    <div class="tasks">
        <div class="task">
            <h1>Week 1</h1>
            <p>Get started by</p>
            <p class="task--sub">task 1</p>
            <p class="task--sub">task 2</p>
        </div>
    </div>
    <div class="btn__box">
        <Link to="/goals-gen">
            <button class="solid">CONTINUE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
