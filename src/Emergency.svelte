<script>
    import { Link } from "svelte-routing";
</script>

<h1>5</h1>
<h2>Calling Emergency Services</h2>
<Link to="/dashboard">
    <h3>CANCEL</h3>
</Link>

<style>
    h3 {
        display: flex;
        justify-content: center;
        padding: 5rem;
        align-items: center;
        border-radius: 50%;
        border: 2px solid var(--color-primary);
        height: 40vw;
        width: 40vw;
        text-align: center;
        position: absolute;
        bottom: 5rem;
        translate: 25%;
    }
</style>
