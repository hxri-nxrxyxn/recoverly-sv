<script>
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";
    import ChatPic from "./assets/undraw_online-posts_avfn.png";
</script>

<main>
    <Nav />
    <h1>Anonymous <span>Chatrooms</span></h1>
    <img src={ChatPic} alt="" class="hero" />
    <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Mollitia,
        laudantium recusandae molestias, odio consequuntur eligendi veniam
        explicabo accusamus error delectus asperiores unde soluta repellendus
        exercitationem facilis quasi perspiciatis enim? Exercitationem.
    </p>
    <div class="chatrooms">
        <Link to="/chat">
            <div class="chatroom chatroom--select">
                <h3>Share my story</h3>
            </div>
        </Link>
        <div class="chatroom">
            <h3>General Recovery</h3>
        </div>
        <div class="chatroom">
            <h3>Tips and Tricks</h3>
        </div>
        <div class="chatroom">
            <h3>Daily Motivation</h3>
        </div>
    </div>
</main>

<style>
    .chatrooms {
        margin-top: 1rem;
    }
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
