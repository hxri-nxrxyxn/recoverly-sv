<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Let it be a new <span>Beginning</span></h1>
    <div class="form__box">
        <div class="form__link">
            <label>Height</label>
            <input type="text" placeholder="CM" required />
        </div>
        <div class="form__link">
            <label>Weight</label>
            <input type="password" placeholder="KG" required />
        </div>
        <div class="form__link">
            <label>Addiction Severity</label>
            <input type="range" min="1" max="100" />
        </div>
        <div class="form__link">
            <label>Existing Medical Conditions</label>
            <input type="text" placeholder="If any" required />
        </div>
    </div>
    <div class="btn__box">
        <div class="progress">
            <div class="progressi">whoa lion!</div>
        </div>
        <Link to="/collect-three">
            <button class="solid">NEXT</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 60%;
    }
</style>
