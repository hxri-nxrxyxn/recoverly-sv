<script>
    import { Link } from "svelte-routing";
    import NavClassic from "./NavClassic.svelte";
</script>

<main>
    <NavClassic />
    <h1>Tell us about <br /><span>Yourself</span></h1>
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
        <div class="progress">
            <div class="progressi">whoa lion!</div>
        </div>
        <Link to="/dashboard">
            <button class="solid">NEXT</button>
        </Link>
    </div>
</main>

<style>
    .option__row {
        justify-content: space-between;
        display: flex;
        gap: 1rem;
        margin-bottom: 1rem;
    }
    .option {
        width: 100%;
        background-color: var(--color-light);
        height: 15vh;
        position: relative;
        border-radius: 1rem;
    }
    .option p {
        position: absolute;
        bottom: 1rem;
        left: 1rem;
    }
    .progressi {
        width: 89%;
    }
    .option:focus {
        background-color: var(--color-primary);
        color: white;
    }
</style>
