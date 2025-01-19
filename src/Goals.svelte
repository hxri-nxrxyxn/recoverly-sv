<script>
    import { Link } from "svelte-routing";
    let duration;
</script>

<main>
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>Goal duration</label>
            <input
                type="text"
                placeholder="Weeks"
                required
                bind:value={duration}
            />
        </div>
    </div>
    <div class="btn__box">
        <Link to={`/goals-gen?duration=${duration}`}>
            <button class="solid">CONTINUE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
