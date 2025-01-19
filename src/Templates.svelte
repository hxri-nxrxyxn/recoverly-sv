<script>
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";
</script>

<main>
    <Nav />
    <h1>Download your <span>Templates</span></h1>
    <div class="template__box">
        <div class="template">
            <h3>All in one Template for schools</h3>
            <p>68 Pages</p>
            <br />
            <button>DOWNLOAD</button>
        </div>
        <div class="template">
            <h3>Template for youth in recovery</h3>
            <p>12 Pages</p>
            <br />
            <button>DOWNLOAD</button>
        </div>
    </div>
    <br />
    <h1>Download your <span>Awareness Guides</span></h1>
    <div class="template__box">
        <div class="template">
            <h3>Comprehensive guide on quitting smoking</h3>
            <p>54 Pages</p>
            <br />
            <button>DOWNLOAD</button>
        </div>
    </div>
</main>

<style>
    .template {
        padding: 1rem;
        border-left: 3px solid var(--color-primary);
        margin-bottom: 1rem;
    }
    .template button {
        border: 0;
        background: var(--color-light);
        padding: 0.25rem 0.5rem;
        border-radius: 0.55rem;
        color: var(--color-text);
    }
</style>
