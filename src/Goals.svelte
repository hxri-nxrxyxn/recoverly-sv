<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>Goal duration</label>
            <input type="text" placeholder="Weeks" required />
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
