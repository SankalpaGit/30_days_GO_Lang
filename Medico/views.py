from django.contrib.auth import get_user_model
from django.shortcuts import get_object_or_404, render, redirect
from django.contrib.auth.models import User
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from .models import Product,Message,RequestBlood, CartItem
from .forms import RequestBloodForm, DonateBloodForm
from django.contrib import messages

# Create your views here.
def Home(request):
    return render(request,'HomePage.html',{}) # redirect to home page

def Signup(request):
    if request.method == 'POST':
        username = request.POST.get('gmail') #get username value from html
        password = request.POST.get('password') #get password value from html
        confirmation = request.POST.get('confirm') #get confirmation
        
        if password == confirmation: #check password and compare 
            # Create a new user and hash the password
            MedicoUser = User.objects.create_user(username=username, password=password) #create new user object
            MedicoUser.save() #save user in admin
            messages.success(request, 'Your account has been created successfully. Please login.')
            return redirect('login')  # Redirect to the login page
        else:
            messages.error(request, "Passwords don't match")  # Error message
            return render(request, 'Register.html', {'error_message': "Passwords don't match"})   
    return render(request, 'Register.html')

def Login(request): #create a login function
    if request.method == 'POST': #method POST
        username = request.POST.get('gmail') 
        password = request.POST.get('password')
        
        # Authenticate the user
        MedicoUser = authenticate(request, username=username, password=password) #check the username and password from object 
        
        if MedicoUser is not None:
            # If user is authenticated, log them in
            login(request, MedicoUser) #object login
            print('Login successful') #print in terminal
            return redirect('Home')  # Redirect to the home page after successful login
        else:
            return render(request, 'Login.html', {'error_message': "Invalid username or password"}) #error message handling
    
    return render(request, 'Login.html')

def Product_List(request):
    Products=Product.objects.all() # call Product models all objects from model.py file
    return render(request,'Product.html', {'products':Products}) #render the called Product object

@login_required
def chat_view(request): #chat function
    if request.method == 'POST':
        content = request.POST.get('content')
        admin_user = get_user_model().objects.get(is_superuser=True)
        # Assuming MedicoUser is the custom user model and the sender of the message
        Message.objects.create(sender=request.user, recipient=admin_user, content=content)
        # Redirect to chat view after sending a message
        return redirect('chats')
    
    # Check if there are any messages for the current user
    messages = Message.objects.filter(recipient=request.user)
    
    return render(request, 'chat.html', {'messages': messages})

@login_required
def view_cart(request):
    cart_items = CartItem.objects.filter(user=request.user)
    return render(request, 'cart.html', {'cart_items': cart_items})

@login_required
def add_to_cart(request):
    if request.method == 'POST':
        product_id = request.POST.get('product_id')
        quantity = 1  # Allow only one quantity 

        # Create or update cart item
        cart_item, created = CartItem.objects.get_or_create(user=request.user, product_id=product_id)
        cart_item.quantity += quantity
        cart_item.save()

    return redirect('products')  # Redirect back to the product list

@login_required
def remove_from_cart(request, product_id):
    if request.method == 'POST':
        # Get the cart item
        cart_item = CartItem.objects.get(user=request.user, product_id=product_id)

        # Delete the cart item
        cart_item.delete()

    return redirect('view_cart')  # Redirect back to the cart view

@login_required
def clear_cart(request):
    if request.method == 'POST':
        # Clear the entire cart
        CartItem.objects.filter(user=request.user).delete()

    return redirect('view_cart')  # Redirect back to the cart view


def About(request):
    return render(request, 'AboutMedico.html',{})

def Blood(request):
    return render(request, 'BloodHome.html',{})

@login_required
def DonateBlood(request):
    # Assuming you fetch donation_id from somewhere
    donation_id = 123  # Replace 123 with the actual donation ID

    if request.method == 'POST':
        form = DonateBloodForm(request.POST)
        if form.is_valid():
            form.save()
            # Redirect to a different URL or view after successful form submission
            return redirect(Blood)  # Replace 'Home' with the name of your success page URL
    else:
        form = DonateBloodForm()
    
    return render(request, 'DonateBlood.html', {'form': form, 'id': donation_id})


@login_required
def RequestBlood(request):
    if request.method=='POST':
        form=RequestBloodForm(request.POST)
        if form.is_valid():
            form.save()
            messages.success(request,'Successfully submitted ! we will contact you soon')
            return redirect(RequestBlood)
        else:
            form=RequestBloodForm()
    return render(request, 'RequestBlood.html',{})



@login_required
def Profile(request):
    return render(request, 'Profile.html', {})

def edit_profile(request):
    return render(request, 'edit_profile.html',{})


def product_search(request):
    query = request.GET.get('q')
    if query:
        results = Product.objects.filter(name__icontains=query)
    else:
        results = None
    return render(request, 'search_results.html', {'results': results, 'query': query})