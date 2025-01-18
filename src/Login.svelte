<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Let's beat our <span>addictions</span></h1>
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
        <Link to="/org-login">
            <button class="outlined">ORGANIZER LOGIN</button>
        </Link>
        <Link>
            <button class="solid">LOGIN</button>
        </Link>
        <p class="foot">
            Don't have an account?<Link to="/register"><b>Create one</b></Link>
        </p>
    </div>
</main>

<style>
</style>
