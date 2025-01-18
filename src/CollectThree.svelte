<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Tell us about <span>Yourself</span></h1>
    <div class="option__box">
        <div class="option__row">
            <div class="option">
                <p>Smoking</p>
            </div>
            <div class="option">
                <p>Alcohol</p>
            </div>
        </div>
        <div class="option__row">
            <div class="option">
                <p>Opium</p>
            </div>
            <div class="option">
                <p>MDMA</p>
            </div>
        </div>
        <div class="option__row">
            <div class="option">
                <p>Cannabis</p>
            </div>
            <div class="option">
                <p>Heroin</p>
            </div>
        </div>
    </div>
    <div class="btn__box">
        <Link to="/dashboard">
            <button class="solid">NEXT</button>
        </Link>
    </div>
</main>

<style>
    .option__row {
        justify-content: space-between;
        display: flex;
    }
</style>
