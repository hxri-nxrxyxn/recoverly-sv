<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Anonymous <span>Chatrooms</span></h1>

    <Link to="/chat">
        <div class="chatrooms">
            <div class="chatroom chatroom--select">
                <h3>Share my story</h3>
            </div>
        </div>
    </Link>
    <div class="chatrooms">
        <div class="chatroom">
            <h3>General Recovery</h3>
        </div>
    </div>
    <div class="chatrooms">
        <div class="chatroom">
            <h3>Tips and Tricks</h3>
        </div>
    </div>
    <div class="chatrooms">
        <div class="chatroom">
            <h3>Daily Motivation</h3>
        </div>
    </div>
</main>

<style>
    .chatroom {
        background-color: var(--color-light);
        height: 30vh;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 1rem;
        border-radius: 0.5rem;
    }
    .chatroom h3 {
        color: var(--color-text);
    }
    .chatroom--select {
        background-color: var(--color-primary);
    }
    .chatroom--select h3 {
        color: white;
    }
</style>
