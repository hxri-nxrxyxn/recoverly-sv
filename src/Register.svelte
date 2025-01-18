<script>
    import { Link } from "svelte-routing";
    import NavClassic from "./NavClassic.svelte";
</script>

<main>
    <NavClassic />
    <h1>Let it be a new <span>Beginning</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>Email Id</label>
            <input type="text" placeholder="Email" required />
        </div>
        <div class="form__link">
            <label>Password</label>
            <input type="password" placeholder="Password" required />
        </div>
    </div>
    <div class="btn__box">
        <Link to="/terms">
            <button class="solid">CONTINUE</button>
        </Link>
        <p class="foot">Check your email for an <span>OTP</span></p>
    </div>
</main>

<style>
</style>
